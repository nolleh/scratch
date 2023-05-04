import { readable } from 'svelte/store';

// readable call only once for start/stop notifier.
//
export default readable({ x: 0, y: 0 }, (set) => {
  document.addEventListener('mousemove', move);
  function move(event: any) {
    // lastMouseMove.x = event.x;
    // lastMouseMove.y = event.y;
    //
    set({
      x: event.clientX,
      y: event.clientY
    });
    return () => {
      document.removeEventListener('mousemove', move);
    };
  }
});
