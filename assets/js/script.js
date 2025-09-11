const menuToggle = document.getElementById('menu-toggle');
const sideMenu = document.getElementById('side-menu');

menuToggle.addEventListener('click', () => {
  sideMenu.classList.toggle('open');
  console.log("p1")
});

document.addEventListener('click', () => {
  console.log("p2")
  if(!sideMenu.contains(event.target) && event.target != menuToggle && sideMenu.classList.contains('open')){
    sideMenu.classList.toggle('open');
  }
});
