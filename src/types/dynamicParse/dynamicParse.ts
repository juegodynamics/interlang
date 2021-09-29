interface Quantity {
  type: "quantity";
  amount: number;
  unit: string;
}

interface Ratio {
  type: "ratio";
  numerator: Quantity;
  denominator: Quantity;
}

type Measurement = Quantity | Ratio;
type MeasurementList = { measurements: Measurement[] };

const sampleData: MeasurementList = {
  measurements: [
    {
      type: "quantity",
      amount: 10,
      unit: "miligrams",
    },
    {
      type: "ratio",
      numerator: {
        type: "quantity",
        amount: 10,
        unit: "miligrams",
      },
      denominator: {
        type: "quantity",
        amount: 20,
        unit: "mililiters",
      },
    },
  ],
};

const main = () => {
  const asString = JSON.stringify(sampleData);
  console.log(asString);
  const asTyped: Measurement[] = JSON.parse(asString);
  console.log(`Successful parsing: ${asTyped}`);
};

main();
