export class Logger {
  debug(msg: string, ...meta: any[]) {
    console.log(msg, meta);
  }
}
