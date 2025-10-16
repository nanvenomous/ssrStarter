import htmx from "htmx.org";
import { dialogEventHandler } from "./dialog";

const w = window as any;
w.htmx = htmx;

w.dialogEventHandler = dialogEventHandler
