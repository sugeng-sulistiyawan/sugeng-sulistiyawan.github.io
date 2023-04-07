const mode = document.getElementById('mode');

if (mode !== null) {

  window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', event => {

    if (event.matches) {

      localStorage.setItem('theme', 'dark');
      document.documentElement.setAttribute('data-dark-mode', '');

    } else {

      localStorage.setItem('theme', 'light');
      document.documentElement.removeAttribute('data-dark-mode');

    }

  })

  mode.addEventListener('click', () => {

    document.documentElement.toggleAttribute('data-dark-mode');
    localStorage.setItem('theme', document.documentElement.hasAttribute('data-dark-mode') ? 'dark' : 'light');

    __updateImageMode();
  });

  if (localStorage.getItem('theme') === 'dark') {

    document.documentElement.setAttribute('data-dark-mode', '');

  } else {

    document.documentElement.removeAttribute('data-dark-mode');

  }

  __updateImageMode();

}

function __updateImageMode() {
  let themesL = document.querySelectorAll('.img-show-light');
  let themesD = document.querySelectorAll('.img-show-dark');
  let isDark = localStorage.getItem('theme') === 'dark';
  for (let i = 0; i < themesL.length; i++) {
    let elemL = themesL[i];
    if (isDark) {
      elemL.style.display = 'none';
    } else {
      elemL.style.display = '';
    }
  }
  for (let i = 0; i < themesD.length; i++) {
    let elemD = themesD[i];
    if (isDark) {
      elemD.style.display = '';
    } else {
      elemD.style.display = 'none';
    }
  }
}
