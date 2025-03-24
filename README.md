# Valtracker

Valtracker is a cli script that displays information about your matches in Valorant using the following api 

https://github.com/Henrik-3/unofficial-valorant-api.

It is completely built in go with the help of the bubbletea framework and go-valorant-api


https://github.com/user-attachments/assets/c6ecfe2d-8f81-4985-b816-88eeb2ca4c2e

## Usage
You will need a api key to use this program. You can get one from this ![discord server](https://discord.com/invite/X3GaVkX2YN).

You can input the key when you launch the program.

![image](https://github.com/user-attachments/assets/12a16f5b-e775-418d-b3c3-5c0ebbd762aa)

## **Build Instructions (Windows)**

### **Prerequisites**
Ensure you have the following installed:
- [Go](https://go.dev/dl/)

### **Clone the Repository**
```sh
git clone https://github.com/Heribio/ValTracker.git
cd ValTracker
```

### **Install Dependencies**
```sh
go mod tidy
```

### **Build the Program**
To compile the binary:
```sh
go build -o valtracker.exe
```

### **Run the Program**
After building, execute:
```sh
valtracker
```

or directly run without building:
```sh
go run .
```

---

## References
https://github.com/charmbracelet/bubbletea

https://github.com/yldshv/go-valorant-api
