# CMSC701_HW2: Bitvector 

This package implements bitvector in Go.  
<br/>


This code is available at https://github.com/lwhite17/CMSC701_HW2.  
<br/>


By Leah White   
<br/><br/>


---------------------------------------------  
---------------------------------------------  
<br/>

# bitvector  

<br/>

This file describes the structure bitvector.BitVector and the bitvector package methods.   

<br/><br/>

---------------------------------------------  
<br/>  

## Structure: bitvector.BitVector  
<br/>

The bitvector.BitVector structure has the following fields:    
<br/>

- *data* ([]byte): The array of uint8 0's and 1's that comprise the bitvector  
<br/>

- *length* (uint64): The length of the bitvector.  

<br/><br/>

---------------------------------------------  
<br/>

## bitvector Methods
<br/>

The package bitvector has the following methods:  
<br/>

- **NewBitVector**(*data* []byte, *length* uint64) *BitVector  
   
    Creates new bit vector.  
    <br/>

    - Parameters:  
        - *data* ([]byte): The array of uint8 0's and 1's that comprise the bitvector  

        - *length* (uint64): The length of the bitvector.  
    <br/>

    - Returns:  
        - \*BitVector: Pointer to the created bitvector.BitVector struct created over *data* with length *length*.  
    <br/>

    - Example call:  

        data := []byte{0, 0, 1, 0, 0, 0, 0, 1, 0, 1}
        var length uint64 = uint64(len(data))
        pointer_to_bitvector := bitvector.NewBitVector(data, length)  
<br/>

- (*b* *BitVector) **Bytes**() []byte  
  
    Gets the bitvector data.  
    <br/>

    - No parameters; method on bitvector.BitVector pointer.  
    <br/>

    - Returns: 
        - []data: The array of uint8 0's and 1's that comprise the bitvector   
    <br/>

    - Example call:  

            data := pointer_to_bitvector.Bytes()
<br/>

- (*b* *BitVector) **Length**() uint64  
   
    Gets the length of the bitvector.  
    <br/>

    - No parameters; method on bitvector.BitVector pointer.  
    <br/>


    - Returns: 
        - uint64: The length of the bitvector.  
    <br/>

    - Example call: 

            length := pointer_to_bitvector.Length()  
<br/>

- (*b* *BitVector) **Get**(*i* int) byte   
   
    Gets the bit at index *i*.  
    <br/>

     - Parameters:  
        - *i* (int): The index at which to retreive the bitvector value.  
    <br/>

    - Returns: 
        - byte: 0 if bitvector is 0 at index *i*; 1 if bitvector is 1 at index *i*.  
    <br/>

    - Example call: 

            var i int = 5
            val := pointer_to_bitvector.Get(i)  
<br/>

- (*b* *BitVector) **GetSlice**(*i* int, *j* int) []byte   
    
    Gets bytes [i, j) from bitvector.  
    <br/>

    - Parameters:  
        - *i* (int): Starting index of slice (inclusive).  

        - *j* (int): Ending index of slice (noninclusive).  
    <br/>

    - Returns: 
        - []byte: Slice of bytes containing bitvector data from index *i* to index *j*.  
    <br/>

    - Example call:  

            var i int = 2
            var j int = 5
            subvector := pointer_to_bitvector.GetSlice(i, j)
<br/>

- (*b* *BitVector) **Save**(*filename* string)   
   
    Saves bitvector structure to file.  
    <br/>

    - Parameters: 
        - *filename* (string): Name of text file in which to save the BitVector structure.  
    <br/>

    - Does not return anything.  
    <br/>

    - Example call: 

            filename := "saved_bitvector.txt"
            pointer_to_bitvector.Save(filename)
<br/>

- **Load**(*filename* string) *BitVector  
    
    Loads bitvector structure from file (made via save method).  
    <br/>

    - Parameters: 
        - *filename* (string): Name of text file from which to load the BitVector structure.  
    <br/>

    - Returns: 
        - *BitVector: BitVector structure loaded from file.  
    <br/>

    - Example call: 

            pointer_to_bitvector := bitvector.Load(filename)

<br/>

- (*b* *BitVector) **Set**(ind uint64, val byte) *BitVector  
    
    Sets the bitvector at index ind to value val.  
    <br/>

    - Parameters: 
        - *ind* (uint64): The index at which to change the bitvector.  

        - *val* (byte): The value to which to change the bitvector.  
    <br/>

    - Returns: 
        - *BitVector: Pointer to updated BitVector.  
    <br/>

    - Example call: 

            ind := 2
            val := 1
            pointer_to_bitvector = pointer_to_bitvector.Set(ind, val)
<br/>


---------------------------------------------   
---------------------------------------------  
