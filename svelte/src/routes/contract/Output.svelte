<script lang="ts">
import { valueStore, store1 } from '../../libs/store2';

  setInterval(() => {
    for (const subscriber of subscribers) {
      subscriber(new Date());
    }
  }, 1000);

  const subscribers: any[] = [];
  const store3 = {
    subscribe(fn: any) {
      // fn('');
      subscribers.push(fn);
      return () => {
        subscribers.splice(subscribers.indexOf(fn), 1);
      };
    },
    set(value:string) {
      for (const subscriber of subscribers) {
        subscriber(value);
      }
    }
  };
</script>

<h1>store1: {$store1}</h1>
<h1>store2: {$valueStore}</h1>

{$store3}

<button on:click= {() => {
  $store3 = 'hello'
}}>click me </button>
