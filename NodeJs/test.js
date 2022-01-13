function calculate(numbers) {
  const postNum = [];
  const negNum = [];
  for(let i = 0; i < numbers.length; i++) {
    if(numbers[i] + 1 === numbers[i+1]) {
      postNum.pop();
      postNum.push(numbers[i]);
      postNum.push(numbers[i+1]);
    }
    if(i + 1 < numbers.length) {
      if(numbers[i] - 1 === numbers[i+1]) {
        negNum.pop();
        negNum.push(numbers[i]);
        negNum.push(numbers[i+1]);
      }
    }
  }
  const unionNum = postNum.filter(element => negNum.includes(element));
  return Math.max(...unionNum);
}

const num3 = [7, 1, 2, 9, 7, 2, 1]
console.log(calculate(num1))