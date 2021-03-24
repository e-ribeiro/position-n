/**
 *  Takes a interger number and returns 
 *  its equivalent based on the rule.
 * @param pos a interger number
 * @returns an equivalent number based on the rule
 */
const findNext = (pos: number): number => {
  const rule: number = pos % 2 == 0 ? pos / 2 : (pos - 1) / 2;
  const number: number = pos * 2 - rule;
  return number;
};

findNext(11); // Output: 17