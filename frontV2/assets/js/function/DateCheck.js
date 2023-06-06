export const DateCheck = (Entrydate) => {
    /*
    * The user must have more than 13 years
    * The date must be after 1900
     */
    const actualDate = new Date() ; 
    const days = actualDate.getDate().toString()
    const months = (actualDate.getMonth() + 1).toString()
    const years = (actualDate.getFullYear() - 13 ).toString()
    if(Entrydate[0]+Entrydate[1]+Entrydate[2]+Entrydate[3] < years && Entrydate[0]+Entrydate[1]+Entrydate[2]+Entrydate[3] > "1900") {
        return true
    } else if(Entrydate[0]+Entrydate[1]+Entrydate[2]+Entrydate[3] < "1900") {
        return "you can't enter a date before 1900 :)"
    } else if(Entrydate[0]+Entrydate[1]+Entrydate[2]+Entrydate[3] > years){
        return "if you don't have 13 years old , you can't create a account"
    } else if(Entrydate[5]+Entrydate[6] > months) {
        return "if you don't have 13 years old , you can't create a account"
    } else if(Entrydate[8]+Entrydate[9] > days) {
        return "if you don't have 13 years old , you can't create a account"
    }
    return true
}

