const delayCallback = (callback: () => void, wait: number) => async () => {
  await new Promise((resolve) => setTimeout(resolve, wait));
  callback();
};

const printOneDelay = delayCallback(() => {
  console.log("One");
}, 1000);

printOneDelay();
console.log("Two");
