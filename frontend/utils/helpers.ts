export const createKey = () : string => {
    let text = "";
    const possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";

    for (var i = 0; i < 20; i++)
        text += possible.charAt(Math.floor(Math.random() * possible.length));

    return text;
}

export const unixToDate = (unixTimeStamp : number) : string =>{
    return new Date(unixTimeStamp).toLocaleString()
}

export const currencyIDR = (number : number) : string => {
    return Intl.NumberFormat("id-ID",{style: "currency", currency:"IDR", minimumFractionDigits:0}).format(number)
}