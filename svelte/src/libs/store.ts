export let value = '';

export function update(newValue: string) {
  value = newValue;
  subscribers.forEach(fn => {
    fn();
  });
}

let subscribers:any[] = [];

export function subscribe(fn: any) {
  subscribers.push(fn);

  return function unsubscribe(fn: any) {
    subscribers.splice(subscribers.indexOf(fn), 1);
  }
}

