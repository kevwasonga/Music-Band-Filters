
# **Music Band Explorer- Filters**

![Music Band Explorer Screenshot](https://via.placeholder.com/800x400.png?text=Music+Band+Explorer+Screenshot)  
*(Replace with an actual screenshot of your project once it's ready)*

**Music Band Explorer** is a web application that allows users to explore and filter music bands based on specific criteria. Users can filter bands by formation year, first album year, number of members, and concert locations. The application features a clean and intuitive interface, making it easy to discover bands that match your preferences.

---

## **Features**

- **Filter Bands by:**
  - Formation Year Range
  - First Album Year Range
  - Number of Members
  - Concert Locations
- **Dynamic Results:** Filtered bands are displayed in real-time.
- **User-Friendly Interface:** Simple and intuitive design.
- **Efficient Backend:** Built with Go (Golang) for fast and reliable performance.
- **Responsive Design:** Works seamlessly on desktop and mobile devices.

---

## **Tech Stack**

- **Backend:** Go (Golang)
- **Frontend:** HTML, CSS, JavaScript
- **Data Storage:** JSON file (`bands.json`)
- **Concurrency:** Go routines for efficient filtering
- **Error Handling:** Graceful error handling and fallbacks

---

## **Project Structure**

```
Music-Band-Filters/
├── data/                    # Contains JSON data file
│   └── bands.json           # JSON file with band data
├── models/                  # Contains all handlers
│   └── ArtistsHandler.go    # Handlers for artist-related requests
├── services/                # Contains JSON handling and manipulation
│   └── unmarshaljson.go     # Functions to load and process JSON data
├── template/                # Contains frontend files (HTML, CSS, JS)
│   ├── index.html           # Main HTML file
│   ├── styles.css           # CSS for styling
│   └── script.js            # JavaScript for dynamic filtering
├── go.mod                   # Go module file
└── main.go                  # Main Go server file
```

---

## **Setup Instructions**

### **Prerequisites**
- Go (Golang) installed on your machine.
- A modern web browser (e.g., Chrome, Firefox).

### **Steps to Run the Project**

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/your-username/music-band-explorer.git
   cd music-band-explorer
   ```

2. **Run the Backend:**
   - Navigate to the `backend` directory:
     ```bash
     cd backend
     ```
   - Start the Go server:
     ```bash
     go run main.go
     ```
   - The backend will start at `http://localhost:8080`.

3. **Open the Frontend:**
   - Open the `frontend/index.html` file in your browser.
   - Alternatively, serve the frontend using a local server (e.g., `python -m http.server` in the `frontend` directory).

4. **Use the Application:**
   - Enter filter criteria (e.g., formation year, concert locations).
   - Click "Apply Filters" to see the filtered bands.

---

## **Usage**

### **Filters**
1. **Formation Year Range:**
   - Enter the minimum and maximum formation year to filter bands.

2. **First Album Year Range:**
   - Enter the minimum and maximum first album year to filter bands.

3. **Number of Members:**
   - Enter the minimum and maximum number of members to filter bands.

4. **Concert Locations:**
   - Select one or more concert locations to filter bands.

### **Example**
- Filter bands formed between **1990** and **2000** with **4–6 members** who performed in **New York, USA**.

---

## **Contributing**

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes.
4. Submit a pull request.

---

## **License**

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.

---

## **Acknowledgments**

- Inspired by the **Groupie Tracker** project.
- Built with ❤️ using Go, HTML, CSS, and JavaScript.

---

## **Contact**

For questions or feedback, feel free to reach out:

- **Your Name**  
- **Email:** your-email@example.com  
- **GitHub:** [your-username](https://github.com/kevwasonga)

