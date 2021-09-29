const failureFunc = () => {
  throw new Error("I failed!");
};

try {
  failureFunc();
  console.log("I win!");
} catch (e) {
  console.log(`Caught error: ${e}`);
}
