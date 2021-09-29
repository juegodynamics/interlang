export interface Circle {
  type: "circle";
  centerX: number;
  centerY: number;
  radius: number;
}

export interface Square {
  type: "square";
  minX: number;
  minY: number;
  maxX: number;
  maxY: number;
}

export type Shape = Circle | Square;

export const getArea = (shape: Shape): number => {
  switch (shape.type) {
    case "circle":
      return Math.PI * shape.radius ** 2;
    case "square":
      return (shape.maxX - shape.minX) * (shape.maxY - shape.minY);
  }
};

export const getTotalArea = (shapes: Shape[]): number =>
  shapes.reduce((total, nextShape) => total + getArea(nextShape), 0);

console.log(
  getTotalArea([
    {
      type: "circle",
      centerX: 0,
      centerY: 0,
      radius: 1 / Math.sqrt(Math.PI),
    },
    {
      type: "square",
      minX: 0,
      minY: 0,
      maxX: 1,
      maxY: 1,
    },
  ])
);
