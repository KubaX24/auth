document.getElementById("c1").focus()

let prevVal = new Map();
document.querySelectorAll('input').forEach((item) => {
    item.addEventListener('input', function(e){
        if(this.checkValidity()){
            prevVal.set(this.id, this.value)
        } else {
            this.value = prevVal.get(this.id);
        }
    });
})

for (let i = 1; i <= 6; i++) {
    document.getElementById("c" + i).addEventListener("keyup", function (e){
        if (e.key !== "Backspace" && e.key !== "ArrowLeft"){
            if (i !== 6)
                document.getElementById("c" + (i+1)).focus()
            else
                document.getElementById("submit").focus()
        }
    })
    document.getElementById("c" + i).addEventListener("keydown", function (e){
        if (e.key === "Backspace"){
            document.getElementById("c" + i).value = ""
            if (i !== 1)
                document.getElementById("c" + (i-1)).focus()
        }else if (e.key === "ArrowLeft"){
            if (i !== 1)
                document.getElementById("c" + (i-1)).focus()
        }
    })
}