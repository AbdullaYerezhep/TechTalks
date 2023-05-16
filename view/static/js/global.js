// Authentication logic

document.querySelectorAll(".signInButton").forEach(signIn =>{
    signIn.addEventListener("click", function() {
        toggleContainerVisibility("signInContainer");
    });
})

document.querySelectorAll(".signUpButton").forEach(signUp => {
    signUp.addEventListener("click", function() {
        toggleContainerVisibility("signUpContainer");
    });
})


let signInContainer = document.getElementById("signInContainer")
let signUpContainer = document.getElementById("signUpContainer")

const overlay = document.getElementById('overlay');
function toggleContainerVisibility(containerId) {
    var container = document.getElementById(containerId);
    if (container.classList.contains("hidden") &&  container === signInContainer)  {
        container.classList.remove("hidden");
        signUpContainer.classList.add("hidden");
        overlay.style.display ="block"
    } else if (container.classList.contains("hidden") && container === signUpContainer) {
        container.classList.remove("hidden");
        signInContainer.classList.add("hidden");
        overlay.style.display ="block"
    }
}

const xmarks = document.querySelectorAll(".fa-x")
xmarks.forEach((x)=>{
    x.addEventListener("click", () => {
        let closingElement = document.getElementById(x.classList.item(x.classList.length -1))
        closingElement.classList.add("hidden")
        overlay.style.display = "none"
    })

})

