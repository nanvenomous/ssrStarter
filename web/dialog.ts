

export function dialogClose(dialogID: string) {
  const dialog = document.getElementById(dialogID) as HTMLDialogElement;
  if (dialog) {
    dialog.close();
  }
}

export function dialogEventHandler (dialogID: string) {
  const dialog = document.getElementById(dialogID) as HTMLDialogElement;
  if (dialog) {
    dialog.addEventListener('click', (event: MouseEvent) => {
      if (event.target instanceof Element && event.target.id === dialogID) {
        dialog.close();
      }
    });

    document.addEventListener('keydown', (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        dialog.close();
      }
    });
  } else {
    console.error("could not find dialog element with id: ", dialogID)
  }
}
