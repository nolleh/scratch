import { writable, readable } from 'svelte/store';

export const valueStore = writable('');

export const store1 = readable(0, (set) => {
  let i = 0;
  const timeoutid = setInterval(() => {
    set(++i);
  }, 1000);

  return () => {
    clearInterval(timeoutid);
  };
});
