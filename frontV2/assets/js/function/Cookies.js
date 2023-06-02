
export const Create_cookie = (name, value) => {
    const expirationDate = new Date();
    expirationDate.setDate(expirationDate.getDate() + 365);
    const toAdd = 'expires='+ expirationDate.toUTCString() +';path=/';
    const cookie = [name, '=', JSON.stringify(value),"; SameSite=Strict;",toAdd].join('');
    document.cookie = cookie;
}

export const get_token = () => {
    let res = String(document.cookie.match(/(?<=[t][o][k][e][n][=])(.)*/))
    res = res.substring(0, res.length - 2)
    return res
}