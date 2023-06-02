export const Footer = () => {
    const linkcss = document.createElement('link'); linkcss.rel = 'stylesheet'; linkcss.href = 'http://localhost:8080/assets/css/footer.css';
    document.head.appendChild(linkcss);
    const section = document.createElement('section');    section.classList.add('footer-section');
    const rights = document.createElement('p');rights.innerHTML = '© 2023 - All rights reserved';
    const p = document.createElement('p');p.innerHTML = 'Folow us';
    const linksdiv = document.createElement('div');linksdiv.classList.add('footer-links');
    const afacebook = document.createElement('a');afacebook.href = 'https://www.facebook.com/';
    const imgfacebook = document.createElement('img');imgfacebook.src = 'http://localhost:8080/assets/img/facebookLogo.svg'; afacebook.appendChild(imgfacebook);

    const ainstagram = document.createElement('a');ainstagram.href = 'https://www.instagram.com/';
    const imginstagram = document.createElement('img');imginstagram.src = 'http://localhost:8080/assets/img/instagramLogo.svg'; ainstagram.appendChild(imginstagram);

    const atwitter = document.createElement('a');atwitter.href = 'https://www.twitter.com/';
    const imgtwitter = document.createElement('img');imgtwitter.src = 'http://localhost:8080/assets/img/twitterLogo.svg'; atwitter.appendChild(imgtwitter);

    const aytb = document.createElement('a');aytb.href = 'https://www.youtube.com/';
    const imgytb = document.createElement('img');imgytb.src = 'http://localhost:8080/assets/img/ytbLogo.svg'; aytb.appendChild(imgytb);
    const aemail = document.createElement('a');aemail.href = 'gurvan.nicolas007@icloud.com';
    const imgemail = document.createElement('img');imgemail.src = 'http://localhost:8080/assets/img/email.svg'; aemail.appendChild(imgemail);
    linksdiv.appendChild(afacebook); linksdiv.appendChild(ainstagram); linksdiv.appendChild(atwitter); linksdiv.appendChild(aytb); linksdiv.appendChild(aemail);
    section.appendChild(p);section.appendChild(linksdiv); section.appendChild(rights);
    document.body.appendChild(section);
}