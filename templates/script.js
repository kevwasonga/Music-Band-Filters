// // Fetch and display bands when the page loads
// document.addEventListener("DOMContentLoaded", function () {
//     fetchBands();
// });

// // Function to fetch bands from the backend
// function fetchBands() {
//     fetch("/bands")
//         .then((response) => response.json())
//         .then((data) => {
//             // Populate concert locations checkboxes
//             populateLocations(data);
//             // Display all bands initially
//             displayBands(data);
//         })
//         .catch((error) => console.error("Error fetching bands:", error));
// }

// // Function to populate concert locations checkboxes
// function populateLocations(bands) {
//     const locationCheckboxes = document.getElementById("locationCheckboxes");
//     const locations = new Set();

//     // Extract unique locations from all bands
//     bands.forEach((band) => {
//         band.concertLocations.forEach((location) => locations.add(location));
//     });

//     // Create checkboxes for each location
//     locations.forEach((location) => {
//         const label = document.createElement("label");
//         const checkbox = document.createElement("input");
//         checkbox.type = "checkbox";
//         checkbox.value = location;
//         label.appendChild(checkbox);
//         label.appendChild(document.createTextNode(location));
//         locationCheckboxes.appendChild(label);
//     });
// }

// // Function to display bands in the results section
// function displayBands(bands) {
//     const bandList = document.getElementById("bandList");
//     bandList.innerHTML = ""; // Clear previous results

//     bands.forEach((band) => {
//         const li = document.createElement("li");
//         li.textContent = `${band.name} (Formed: ${band.formationYear}, Members: ${band.members})`;
//         bandList.appendChild(li);
//     });
// }

// // Event listener for the filter button
// document.getElementById("filterButton").addEventListener("click", function () {
//     const filters = {
//         minFormationYear: document.getElementById("minFormationYear").value,
//         maxFormationYear: document.getElementById("maxFormationYear").value,
//         minFirstAlbumYear: document.getElementById("minFirstAlbumYear").value,
//         maxFirstAlbumYear: document.getElementById("maxFirstAlbumYear").value,
//         minMembers: document.getElementById("minMembers").value,
//         maxMembers: document.getElementById("maxMembers").value,
//         locations: Array.from(document.querySelectorAll("#locationCheckboxes input:checked")).map(
//             (checkbox) => checkbox.value
//         ),
//     };

//     // Send filter request to the backend
//     fetch("/filter", {
//         method: "POST",
//         headers: {
//             "Content-Type": "application/json",
//         },
//         body: JSON.stringify(filters),
//     })
//         .then((response) => response.json())
//         .then((data) => displayBands(data))
//         .catch((error) => console.error("Error applying filters:", error));
// });