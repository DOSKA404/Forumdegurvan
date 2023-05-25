import { Login } from "./Components/login.js";
import { Footer } from "./Components/footer.js";
import {menu} from "./Components/mainPage.js";

let linksCss = []

const LoginPage = () => {
    Login();
    Footer();
}
const MainPage =()=> {
    menu()
    Footer()
}
MainPage()