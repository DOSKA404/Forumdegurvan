export const Footer = () => {
    const linkcss = document.createElement('link'); linkcss.rel = 'stylesheet';linkcss.href = 'http://localhost:8080/assets/css/footer.css';
    document.head.appendChild(linkcss);
    const section = document.createElement('section');
    section.classList.add('footer-section');
    document.body.appendChild(section);

}