<template>
  <div class="calculator">
    <div class="display">{{ display }}</div>
    <div class="buttons">
      <button @click="appendDigit(7)">7</button>
      <button @click="appendDigit(8)">8</button>
      <button @click="appendDigit(9)">9</button>
      <button @click="appendOperator('+')">+</button>
      <button @click="appendDigit(4)">4</button>
      <button @click="appendDigit(5)">5</button>
      <button @click="appendDigit(6)">6</button>
      <button @click="appendOperator('-')">-</button>
      <button @click="appendDigit(1)">1</button>
      <button @click="appendDigit(2)">2</button>
      <button @click="appendDigit(3)">3</button>
      <button @click="appendOperator('*')">*</button>
      <button @click="appendDigit(0)">0</button>
      <button @click="appendDecimal()">.</button>
      <button @click="calculate">=
      </button>
      <button @click="appendOperator('/')">/</button>
      <button @click="clear">C</button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CalculatorPage',
  data() {
    return {
      display: '0',
      currentNumber: '',
      operator: null,
      previousNumber: null
    };
  },
  methods: {
    appendDigit(digit) {
      if (this.display === '0') {
        this.display = digit.toString();
      } else {
        this.display += digit.toString();
      }
    },
    appendDecimal() {
      if (!this.display.includes('.')) {
        this.display += '.';
      }
    },
    appendOperator(operator) {
      if (this.operator !== null) {
        this.calculate();
      }
      this.operator = operator;
      this.previousNumber = parseFloat(this.display);
      this.currentNumber = '';
    },
    clear() {
      this.display = '0';
      this.operator = null;
      this.previousNumber = null;
    },
    calculate() {
      const current = parseFloat(this.display);
      let result;
      switch (this.operator) {
        case '+':
          result = this.previousNumber + current;
          break;
        case '-':
          result = this.previousNumber - current;
          break;
        case '*':
          result = this.previousNumber * current;
          break;
        case '/':
          result = this.previousNumber / current;
          break;
        default:
          return;
      }
      this.display = result.toString();
      this.operator = null;
      this.previousNumber = null;
    }
  }
};
</script>

<style scoped>
.calculator {
  width: 200px;
  margin: 0 auto;
}

.display {
  background-color: #f0f0f0;
  padding: 10px;
  text-align: right;
}

.buttons {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 5px;
}

button {
  padding: 10px;
}
</style>
