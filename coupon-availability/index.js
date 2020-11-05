const express = require('express');
const app = express();

app.use(express.urlencoded());
app.use(express.json());

app.post('/', (_, res) => {
  return res.json({ Status: getValidInvalid() });
});

app.listen(9093, () => {});

function getValidInvalid() {
  const randomNumber = Math.floor(Math.random() * 10);
  return isEven(randomNumber) ? 'valid' : 'invalid';
}

function isEven(number) {
  return number % 2 === 0 ? true : false;
}