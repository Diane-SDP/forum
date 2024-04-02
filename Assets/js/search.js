const searchContainerEl = document.querySelector(".search-container");
const closeBtn = document.querySelector(".fa-times");

searchContainerEl.addEventListener("click", () => {
  searchContainerEl.classList.add("active");
});

closeBtn.addEventListener("click", (event) => {
    event.stopPropagation();
    searchContainerEl.classList.remove("active");
    search = document.querySelector('input').value = "";
});



