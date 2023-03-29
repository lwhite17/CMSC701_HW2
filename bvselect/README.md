# CMSC701_HW2: Bitvector 

This package implements log-time bitvector select in Go.  
<br/>


This code is available at https://github.com/lwhite17/CMSC701_HW2.  
<br/>


By Leah White   
<br/><br/>


---------------------------------------------  
---------------------------------------------  
<br/>

# bitvector select: bvselect  

<br/>

This file describes the structure bvselect.Select and the bvselect package methods.   

<br/><br/>

---------------------------------------------  
<br/>  

## Structure: bvselect.Select  
<br/>

The bvselect.Select structure has the following fields:    
<br/>

- *b* (*bitvector.BitVector): Pointer to the underlying BitVector struct.  
<br/>

- *r* (*bvrank.Rank*): Pointer to the undrelying Rank struct (which is over the same underlying bitvector).  

<br/><br/>

---------------------------------------------  
<br/>

## bvselect Methods
<br/>

The package bvselect has the following methods:  
<br/>

- **NewSelectStruct**(*b* *bitvector.BitVector, *r* *bvrank.Rank) *Select  
   
    Creates new Select struct over *b* and *r*.  
    <br/>

    - Parameters:  
        - *b* (*bitvector.BitVector): The underlying bitvector.BitVector struct.

        - *r* (*bvrank.Rank): The underlying bvrank.Rank struct (made over the same bitvector).  
    <br/>

    - Returns:  
        - \*Select: Pointer to the created bvselect.Select struct created over *b* and *r*.  
    <br/>

    - Example call:  

        data := []byte{0, 0, 1, 0, 0, 0, 0, 1, 0, 1}
        var length uint64 = uint64(len(data))
        b := bitvector.NewBitVector(data, length)  
        r := bvrank.NewRankStruct(b)
        s := bvselect.NewSelectStruct(b, r)  
<br/>

- (*s* *Select) **Select1**(*i* int) uint64  

    Calls log-time Select1 using binary search. Gets the first index of the bitvector where the rank is *i*.  
    <br/>

    - Parameters:  

        - *i* (int): The rank value to conduct select on. 
    <br/>

    - Returns: 
        - uint64: The index of the bitvector of the first index for which Rank1(index) = *i*.
    <br/>

    - Example call:  

            var i uint64 = 2
            ind := s.Select1(i)  
<br/>

- (*s* *Select) **Overhead**() uint64  
   
    Gets the overhead of the current Select struct. Note that since this implementation is not Clark's select and since the Select struct points to the Rank structure and does not store it twice, this overhead is essentially the same as that of the bvrank.Rank structure. 
    <br/>

    - No parameters; method on bvselect.Select pointer.  

    <br/>

    - Returns: 
        - uint64: The overhead (amount of bits) needed to store the current Select structure.  
    <br/>

    - Example call: 

            overheadbits := s.Overhead()  
<br/>

- (*s* \*Select) **Save**(*filename* string)   
   
    Saves Select structure to file. Since the select structure stores no additional information to the Rank structure, this simply saves out the underlying Rank structure (and hence the underlying BitVector struct as well).  
    <br/>

    - Parameters: 
        - *filename* (string): Name of text file in which to save the BitVector structure.  
    <br/>

    - Does not return anything.  
    <br/>

    - Example call: 

            filename := "saved_select.txt"
            s.Save(filename)
<br/>

- **Load**(*filename* string) *Select  
    
    Loads the select structure from file made via save method. Inherently loads the underlying Rank and BitVector structures as well.  
    <br/>

    - Parameters: 
        - *filename* (string): Name of text file from which to load the Select structure.  
    <br/>

    - Returns: 
        - *BitVector: Select structure loaded from files.  
    <br/>

    - Example call: 

            pointer_to_bitvector := bitvector.Load(filename)

<br/>


---------------------------------------------   
---------------------------------------------  
