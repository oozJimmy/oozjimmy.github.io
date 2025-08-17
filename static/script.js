const menuToggle = document.getElementById('menu-toggle');
const sideMenu = document.getElementById('side-menu');
const main = document.querySelector('main');

menuToggle.addEventListener('click', () => {
  sideMenu.classList.toggle('open');
  main.classList.toggle('shifted');
});