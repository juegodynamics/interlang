export class Signal<T> {
  promise: Promise<T>;
  resolve: (value?: T | PromiseLike<T>) => void;
  reject: (reason?: any) => void;

  constructor() {
    this.promise = new Promise((resolve, reject) => {
      this.resolve = resolve;
      this.reject = reject;
    });
  }
}

export class Channel<T> {
  public constructor(
    public readonly capacity = 0,
    private readonly values: Array<T> = [],
    private readonly sends: Array<{ value: T; signal: Signal<void> }> = [],
    private readonly recvs: Array<Signal<T>> = []
  ) {}

  public async send(value: T): Promise<void> {
    if (this.recvs.length > 0) {
      this.recvs.shift().resolve(value);
      return;
    }

    if (this.values.length < this.capacity) {
      this.values.push(value);
      return;
    }

    const signal = new Signal<void>();
    this.sends.push({ value, signal });
    await signal.promise;
  }

  public async recv(): Promise<T> {
    if (this.values.length > 0) return this.values.shift();

    if (this.sends.length > 0) {
      const send = this.sends.shift();
      send.signal.resolve();
      return send.value;
    }

    const signal = new Signal<T>();
    this.recvs.push(signal);
    return await signal.promise;
  }
}

export const fibonacci = async (c: Channel<number>, quit: Channel<number>): Promise<void> => {
  let nums = {
    x: 0,
    y: 1,
  };

  let done = false;

  while (true) {
    await Promise.race([
      c.send(nums.x).then(() => {
        nums = { x: nums.y, y: nums.x + nums.y };
      }),
      quit.recv().then(() => {
        done = true;
        console.log("quit");
      }),
    ]);

    if (done) {
      return;
    }
  }
};

const main = () => {
  const c = new Channel<number>();
  const quit = new Channel<number>();

  (async () => {
    for (let i = 0; i < 10; i++) {
      console.log(await c.recv());
    }
    quit.send(0);
  })();

  fibonacci(c, quit).then(() => {});
};

main();
