
const themeLocalStorageID = "theme"
const themeDefault = "emerald"
const themeDark = "synthwave"
const themeDataAttribute = "data-theme"
const themeControllersSelector = "input.theme-controller"

function updateAllThemeToggles(checked: boolean) {
  const allToggles = document.querySelectorAll(themeControllersSelector);
  allToggles.forEach(toggle => {
    (toggle as HTMLInputElement).checked = checked;
  });
}

export function applySavedTheme() {
  const savedTheme = localStorage.getItem(themeLocalStorageID);
  document.documentElement.setAttribute(themeDataAttribute, savedTheme || themeDefault);
  var checked: boolean
  if (savedTheme === themeDark) {
    checked = true;
  } else {
    checked = false;
  }
  document.addEventListener("DOMContentLoaded", () => {
    updateAllThemeToggles(checked)
    updateThemeDisplay()
  });
}

export function persistTheme(checkbox: HTMLInputElement) {
  const theme = checkbox.checked ? themeDark : themeDefault;
  localStorage.setItem(themeLocalStorageID, theme);
  document.documentElement.setAttribute(themeDataAttribute, theme);
  updateAllThemeToggles(checkbox.checked)
  updateThemeDisplay()
}

export function updateThemeDisplay() {
  const currentThemeElement = document.getElementById('current-theme');
  if (currentThemeElement) {
    const currentTheme = document.documentElement.getAttribute(themeDataAttribute) || themeDefault;
    const displayName = currentTheme === themeDark ? 'Synthwave (Dark)' : 'Emerald (Light)';
    currentThemeElement.textContent = displayName;
  }
}
