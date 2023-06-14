document.getElementById("area").addEventListener("keypress", function (e) {
    if (e.key === 'Enter') {
        document.getElementById("date").value = date(new Date().toLocaleDateString())
        document.getElementById('post').submit();
    }
    
});

document.getElementById("area").addEventListener("input", function (e) {
    document.getElementById("content").value = e.target.value;
});

function date(Date) {
    return Date[6] + Date[7] + Date[8] + Date[9]+"-"+Date[3] + Date[4]+"-"+Date[0] + Date[1];
}

document.getElementById("openAccountOptions").addEventListener("click", () => {
    document.getElementById("AccountOption").classList.toggle("hidden");
    document.getElementById("AccountOption").classList.toggle("AccountOption");
    document.getElementById("openAccountOptions").classList.toggle("margin-left");
});