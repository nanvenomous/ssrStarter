import htmx from "htmx.org";
import { dismissAlert } from './alert'
import { dialogEventHandler } from "./dialog";

const w = window as any;
w.htmx = htmx;

w.dismissAlert = dismissAlert
w.dialogEventHandler = dialogEventHandler
