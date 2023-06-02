export const menu = ()=>{
const linkcss = document.createElement('link');linkcss.rel = 'stylesheet';linkcss.href = 'http://localhost:8080/assets/css/mainPage.css'; linkcss.id = "";document.head.appendChild(linkcss);
const div = document.createElement('div'); div.classList.add('menu-div');
const section = document.createElement('section'); section.classList.add('menu-section');
const navbar = document.createElement('div'); navbar.classList.add('navbar-search');
const burgerMenu = document.createElement('div'); burgerMenu.classList.add('burger-menu');
const  burgerOpen = document.createElement('i'); burgerOpen.classList.add('fas'); burgerOpen.classList.add('fa-bars'); burgerOpen.classList.add('ouvrir');
const  burgerClose = document.createElement('i'); burgerClose.classList.add('fas'); burgerClose.classList.add('fa-times'); burgerClose.classList.add('fermer');
const titlefont = document.createElement('div'); titlefont.classList.add('title-font');
const title = document.createElement('h1'); title.classList.add('CLOAK')
const span1 = document.createElement("span") ; span1.classList.add("first-word");span1.innerText = 'CLOAK' ;title.appendChild(span1)
const span2 = document.createElement("span") ; span1.classList.add("second-word");span2.innerText = '& DAGGER'; title.appendChild(span2)
const searchBar = document.createElement('div'); searchBar.classList.add('search-bar', 'custom-width');
const searchInput = document.createElement('input');  searchInput.classList.add('search-input');   searchInput.setAttribute('type', 'text'); searchInput.setAttribute('placeholder', 'Recherche...');
const searchButton = document.createElement('button'); searchButton.classList.add('search-button'); searchButton.innerText = 'Rechercher';
searchBar.appendChild(searchInput); searchBar.appendChild(searchButton);
burgerMenu.appendChild(burgerOpen); burgerMenu.appendChild(burgerClose)
titlefont.appendChild(title);
div.appendChild(searchBar);
div.appendChild(titlefont); div.appendChild(burgerMenu);
div.appendChild(navbar); 
section.appendChild(div)
document.body.appendChild(section);
}