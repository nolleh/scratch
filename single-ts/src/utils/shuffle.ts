import { randomInt } from "crypto";

export class Shuffle {
  static shuffle<T>(array: Array<T>, start: number, length: number){ 
    const prior = array.slice(0, start);
    const end = array.slice(length);
    const arr = array.slice(start, length);
    for (let i = 0; i < length; i++) {
      const fst = randomInt(length - start);
      const snd = randomInt(length - start);
      const tmp = arr[fst];
      arr[fst] = arr[snd];
      arr[snd] = tmp;
    }
    return prior.concat(arr).concat(end)
  }
}
