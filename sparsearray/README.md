# CMSC701_HW2: Bitvector 

This package implements sparse array in Go.  
<br/>


This code is available at https://github.com/lwhite17/CMSC701_HW2.  
<br/>


By Leah White   
<br/><br/>


---------------------------------------------  
---------------------------------------------  
<br/>

# sparsearray  

<br/>

This file describes the structure sparsearray.SparseArray and the sparsearray package methods.   

<br/><br/>

---------------------------------------------  
<br/>  

## Structure: sparsearray.SparseArray  
<br/>

The sparsearray.SparseArray structure has the following fields:    
<br/>

- *bitv* (*bitvector.BitVector): Pointer to the underlying BitVector struct.  
<br/>

- *rank* (*bvrank.Rank): Pointer to the underlying Rank struct.  
<br/>

- *sel* (*bvselect.Select): Pointer to the underlying Select struct.  
<br/>

- *strings* ([]string): Array (slice) of dense string values.  


<br/><br/>

---------------------------------------------  
<br/>

## sparsearray Methods
<br/>

The package sparsearray has the following methods:  
<br/>

- **Create**(*size* uint64) *SparseArray  
   
    Creates new SparseArray struct.  
    <br/>

    - Parameters:  
        - *size* (uint64): The size of the sparse array to create.  
    <br/>

    - Returns:  
        - \*SparseArray: Pointer to the created sparsearray.SparseArray struct.  
    <br/>

    - Example call:  

        var size uint64 = 100
        sa := sparsearray.Create(size)    
<br/>

- (*sa* *SparseArray) **Append**(*elem* string, *pos* uint64) *SparseArray
  
    Appends the string *elem* to the *pos*th position in the bitvector. This changes the underlying bitvector value at the index *pos* to 1 and inserts the string *elem* into *sa*'s *strings* field.    
    <br/>

    - Parameters:  

        - *elem* (string): String to insert into array.  

        - *pos* (uint64): Position at which to insert *elem*.      
    <br/>

    - Returns: 
        - \*SparseArray: Pointer to the updated sparsearray.SparseArray struct.  
    <br/>

    - Example call:  

            elem := "element1"
            var pos uint64 = 2
            sa := sa.Append(elem, pos)
<br/>

- (*sa* *SparseArray) **Finalize**() *SparseArray  
   
    "Finalizes" the SparseArray struct. At this point, no more elements will be added, so this is when the Rank and Select data structures are created.  ***This function must be called before GetAtRank, GetAtIndex, and GetIndexOf methods are called.***  
    <br/>

    - No parameters; method on sparsearray.SparseArray pointer.  
    <br/>

    - Returns: 
        - \*SparseArray: Pointer to the finalized sparsearray.SparseArray struct.
    <br/>

    - Example call: 

            sa := sa.Finalize()  
<br/>

- (*sa* *SparseArray) **GetAtRank**(*r* uint64) (*string, bool)   
   
    Gets the *r*-th element of the sparse array. Returns true and a pointer to the respective string if there are $\geq$ r items in the sparse array. Otherwise, it returns false and nil as a pointer.  

    <br/>

     - Parameters:  
        - *r* (uint64): Index of the string element to return.   
    <br/>

    - Returns: 
        - \*string: Returns nil or Pointer to *r*-th string.  

        - bool: Returns false (if <> r values in array) or true (if $\geq$ r values in array).  
        
        
    <br/>

    - Example call: 

            var r uint64 = 5
            stringpointer, boolval := sa.GetAtRank(r)    
<br/>

- (*sa* *SparseArray) **GetAtIndex**(*r* uint64) (*string, bool)
    
    Looks at the *r*-th index of the bitvector. If it is 1, then this method returns a pointer to the string corresponding to that index and true. Otherwise, it returns nil and false.  
    <br/>

    - Parameters:  
        - *r* (uint64): Index of the bitvector to query.  

    <br/>

   - Returns: 
        - \*string: Returns nil (if *r*-th element of underlying bitvector is 0) or Pointer to *r*-th string (if *r*-th element of underlying bitvector is 1).  

        - bool: Returns false (if *r*-th element of underlying bitvector is 0) or true (if *r*-th element of underlying bitvector is 1).  
    <br/>

    - Example call:  

            var r uint64 = 42
            stringpointer, boolval := sa.GetAtIndex(r)   
<br/>

- (*sa* *SparseArray) **GetIndexOf**(*r* uint64) uint64
    
    Returns the index in the sparse array where the *r*-th present element appears, using select on the bitvector. If *r* is larger than the number of present elements, then this function returns the max uint64 value.     
    <br/>

    - Parameters:  
        - *r* (uint64): The rank value on which to conduct Select1.  

    <br/>

   - Returns: 
        -  uint64: The index in the sparse array at which the *r*-th present element appears. If *r* is larger than the number of elements in the sparse array, the max Uint64 is returned.  
    <br/>

    - Example call:  

            var r uint64 = 3
            ind := sa.GetIndexOf(r)   
<br/>

- (*sa* *SparseArray) **Size**() uint64
    
    Gets the size of the sparse array 
    <br/>

    - No parameters; method on sparsearray.SparseArray pointer.  
    <br/>

   - Returns: 
        - uint64: The size of the sparse array (equivalently, the length of the underlying bitvector.)   
    <br/>

    - Example call:  

            size := sa.Size()   
<br/>

- (*sa* *SparseArray) **NumElemAt**(*r* uint64) uint64
    
    Returns 1 + the rank of the input index (because exclusive rank was implemented), which is the number of elements up to index *r* (inclusive).    
    <br/>

    - Parameters:  
        - *r* (uint64): Input index at which to find rank.  

    <br/>

   - Returns: 
        - uint64: Rank of the index *r* + 1   
    <br/>

    - Example call:  

            var r uint64 = 42
            rrank := sa.NumElemAt(r)   
<br/>

- (*sa* *SparseArray) **Save**(*filename* string)   
   
    Saves the SparseArray structure to file. Inherently saves the underlying Rank and Bitvector structs to files using their respective save methods. Additionally, saves the array of strings to the amended filename (where the end of it (".txt") is replaced by "_sparsearray.txt"). The Select struct is not saved as it requires no additional information to the Rank struct.  
    <br/>

    - Parameters: 
        - *filename* (string): Name of text file in which to save the BitVector structure.  
    <br/>

    - Does not return anything.  
    <br/>

    - Example call: 

            filename := "example_name.txt"
            sa.Save(filename)  

        Note that the array of strings would be saved to "example_name_sparsearray.txt", the underlying Rank struct would be saved to "example_name_rank.txt", and the underlying BitVector struct would be saved to "example_name_bitvector.txt".  
<br/>

- **Load**(*filename* string) *SparseArray  
    
    Loads SparseArray structure from file (made via save method). Loads the array of strings, Rank struct, and BitVector struct from filename with the end amended to "_sparsearray.txt", "_rank.txt", and "_bitvector.txt" respectively. Creates the Select struct based off the loaded Rank and BitVector structs.  
    <br/>

    - Parameters: 
        - *filename* (string): Name of text file from which to load the BitVector structure.  
    <br/>

    - Returns: 
        - *SparseArray: SparseArray structure loaded from files.  
    <br/>

    - Example call: 

            sa := sparsearray.Load(filename)


<br/><br/>

### Supporting methods (not required by assignment but required to run implementation)  
<br/>

- (*sa* *SparseArray) **GetBytes**() []byte
    
    Gets the underylying bitvector's data.
    <br/>

    - No parameters; method on sparsearray.SparseArray pointer.    
    <br/>

    - Returns: 
        - []byte: Array of 0's and 1's that comprise the underlying bitvector.
    <br/>

    - Example call: 

            data := sa.GetBytes()
<br/>

- (*sa* *SparseArray) **GetStrings**() []strings
    
    Gets the dense array of strings from the SparseArray.
    <br/>

    - No parameters; method on sparsearray.SparseArray pointer.    
    <br/>

    - Returns: 
        - []strings: Dense array of string elements of the sparse array.
    <br/>

    - Example call: 

            strings := sa.GetStrings()
<br/>


---------------------------------------------   
---------------------------------------------  
