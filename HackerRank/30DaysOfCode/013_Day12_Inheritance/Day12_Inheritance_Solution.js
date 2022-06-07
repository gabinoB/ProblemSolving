'use strict';

var _input = '';
var _index = 0;
process.stdin.on('data', (data) => { _input += data; });
process.stdin.on('end', () => {
    _input = _input.split(new RegExp('[ \n]+'));
    main();    
});
function read() { return _input[_index++]; }

/**** Ignore above this line. ****/

class Person {
    constructor(firstName, lastName, identification) {
        this.firstName = firstName;
        this.lastName = lastName;
        this.idNumber = identification;
    }
    
    printPerson() {
        console.log(
            "Name: " + this.lastName + ", " + this.firstName 
            + "\nID: " + this.idNumber
        )
    }
}

class Student extends Person {
    /*    
    *   Class Constructor
    *   
    *   @param firstName - A string denoting the Person's first name.
    *   @param lastName - A string denoting the Person's last name.
    *   @param id - An integer denoting the Person's ID number.
    *   @param scores - An array of integers denoting the Person's test scores.
    */
    // Write your constructor here
    //super(firstName, lastName, id)
    constructor(firstName, lastName, id, testScores) {
        super(firstName, lastName, id);
        this.testScores = testScores;
    }
    
    /*    
    *   Method Name: calculate
    *   @return A character denoting the grade.
    */
    // Write your method here
    calculate() {

        // Page where find the solution to my problem, was i didnt put the ".this" before "testScores.length"
        // https://programs.programmingoneonone.com/2021/01/hackerrank-day-12-inheritance-solution.html
        // Credits to him. :D

        let sum = 0
        let result = 0
        for (var i = 0; i < this.testScores.length; i++) {
            sum += this.testScores[i]
        }
        result = sum / this.testScores.length

        if (result >= 90) {
            return "O"
        } else if (result >= 80) {
            return "E"
        } else if (result >= 70) {
            return "A"
        } else if (result >= 55) {
            return "P"
        } else if (result >= 40) {
            return "D"
        } else {
            return "T"
        }
    }
    
}

function main() {
    let firstName = read()
    let lastName = read()
    let id = +read()
    let numScores = +read()
    let testScores = new Array(numScores)
    
    for (var i = 0; i < numScores; i++) {
        testScores[i] = +read()  
    }

    let s = new Student(firstName, lastName, id, testScores)
    s.printPerson()
    s.calculate()
    console.log('Grade: ' + s.calculate())
}
