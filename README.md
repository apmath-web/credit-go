# credit-go
Credit API on Go lang


    /*
    Just an example how you should implement an interface
    */
    
    type Person struct {
        firstName, lastName string
    }
    
    func (p Person) Person(firstName string, lastName string) {
        p.firstName = firstName
        p.lastName = lastName
    }
    
    func (p Person) GetFirstName() string {
        return p.firstName
    }
    
    func (p Person) GetLastName() string {
        return p.lastName
    }
