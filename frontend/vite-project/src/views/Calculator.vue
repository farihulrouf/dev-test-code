<template>
    <div class="calculator bg-gray-200 p-4 rounded-lg shadow-md w-72 mx-auto">
      <div class="display bg-white p-2 text-right rounded-lg mb-4">{{ display }}</div>
      <div class="grid grid-cols-4 gap-2">
        <button class="btn" @click="appendToDisplay('7')">7</button>
        <button class="btn" @click="appendToDisplay('8')">8</button>
        <button class="btn" @click="appendToDisplay('9')">9</button>
        <button class="btn btn-orange" @click="setOperator('+')">+</button>
  
        <button class="btn" @click="appendToDisplay('4')">4</button>
        <button class="btn" @click="appendToDisplay('5')">5</button>
        <button class="btn" @click="appendToDisplay('6')">6</button>
        <button class="btn btn-orange" @click="setOperator('-')">-</button>
  
        <button class="btn" @click="appendToDisplay('1')">1</button>
        <button class="btn" @click="appendToDisplay('2')">2</button>
        <button class="btn" @click="appendToDisplay('3')">3</button>
        <button class="btn btn-orange" @click="setOperator('*')">*</button>
  
        <button class="btn" @click="clearDisplay()">C</button>
        <button class="btn" @click="appendToDisplay('0')">0</button>
        <button class="btn" @click="appendToDisplay('.')">.</button>
        <button class="btn btn-orange" @click="setOperator('/')">/</button>
  
        <button class="btn btn-equal col-span-2" @click="calculate()">=</button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        display: ''
      };
    },
    methods: {
      appendToDisplay(value) {
        this.display += value;
      },
      setOperator(operator) {
        if (this.display !== '' && isNaN(this.display[this.display.length - 1])) {
          this.display = this.display.slice(0, -1) + operator;
        } else {
          this.display += operator;
        }
      },
      clearDisplay() {
        this.display = '';
      },
      calculate() {
        try {
          this.display = eval(this.display).toString();
        } catch (error) {
          this.display = 'Error';
        }
      }
    }
  };
  </script>
  
  <style>
  .btn {
    background-color: #d1d5db;
    color: #1f2937;
    border: 1px solid #9ca3af;
    padding: 10px;
    border-radius: 5px;
    transition: background-color 0.3s ease;
  }
  
  .btn:hover {
    background-color: #9ca3af;
  }
  
  .btn-orange {
    background-color: #f59e0b;
    color: #ffffff;
  }
  
  .btn-orange:hover {
    background-color: #d97706;
  }
  
  .btn-equal {
    background-color: #4b5563;
    color: #ffffff;
  }
  </style>
  