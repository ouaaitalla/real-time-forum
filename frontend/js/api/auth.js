import { datareg } from "../utils/date.js";
import {datalogin} from "../utils/date.js";
export async function registerapi() {
    const response = await fetch("http://localhost:8080/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(datareg),
    });

    const data = await response.json();
    // console.log(data);
    return data
}
export async function loginapi() {
    const response = await fetch("http://localhost:8080/login",{
        method: "POST",
        headers : {
            "Content-Type": "application/json"
        },
        body : JSON.stringify(datalogin)
    })
    const data = await response.json();
    console.log(data);
    return data
}
export function getloginData(){
    datalogin.nickname = document.getElementById("identifier").value;
    datalogin.password = document.getElementById("password").value;
}
export function getregisterData() {
    datareg.nickname = document.getElementById("nickname").value;
    datareg.firstname = document.getElementById("firstname").value;
    datareg.lastname = document.getElementById("lastname").value;
    datareg.age = document.getElementById("age").value;
    datareg.gender = document.getElementById("gender").value;
    datareg.email = document.getElementById("email").value;
    datareg.password = document.getElementById("password").value;
}