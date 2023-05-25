export const menu = ()=>{
const linkcss = document.createElement('link');linkcss.rel = 'stylesheet';linkcss.href = 'http://localhost:8080/assets/css/mainPage.css'; linkcss.id = "";document.head.appendChild(linkcss);
const section = document.createElement('section'); section.classList.add('menu-section');
const navbar = document.createElement('div'); navbar.classList.add('navbar-search');
const burgerMenu = document.createElement('div'); burgerMenu.classList.add('burger-menu');
const  burgerOpen = document.createElement('i'); burgerOpen.classList.add('fas'); burgerOpen.classList.add('fa-bars'); burgerOpen.classList.add('ouvrir');
const  burgerClose = document.createElement('i'); burgerClose.classList.add('fas'); burgerClose.classList.add('fa-times'); burgerClose.classList.add('fermer');
const titlefont = document.createElement('div'); titlefont.classList.add('title-font');
const title = document.createElement('h1'); title.classList.add('CLOAK-&-DAGGER')
burgerMenu.appendChild(burgerOpen); burgerMenu.appendChild(burgerClose)
titlefont.appendChild(title);
section.appendChild(titlefont); section.appendChild(burgerMenu);
section.appendChild(navbar); 
document.body.appendChild(section);
console.log(section)
}