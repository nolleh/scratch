import { Shuffle } from "utils/shuffle";
import { Logger } from "logger";

new Logger().debug("hello,world");
new Logger().debug("shuffle", Shuffle.shuffle([1, 2, 3, 4, 5], 1, 4));
