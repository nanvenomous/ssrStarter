import htmx from "htmx.org";
import { dismissAlert } from './alert'
import { dialogEventHandler } from "./dialog";
import { applySavedTheme, persistTheme, updateThemeDisplay } from './theme'

const w = window as any;
w.htmx = htmx;

w.dismissAlert = dismissAlert
w.dialogEventHandler = dialogEventHandler
w.applySavedTheme = applySavedTheme
w.persistTheme = persistTheme
w.updateThemeDisplay = updateThemeDisplay
