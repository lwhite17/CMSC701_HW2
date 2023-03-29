# CMSC701_HW2: Bitvector 

This package implements Jacobson's bitvector rank (***exclusive***) in Go.  
<br/>


This code is available at https://github.com/lwhite17/CMSC701_HW2.  
<br/>


By Leah White   
<br/><br/>


---------------------------------------------  
---------------------------------------------  
<br/>

# bitvector rank: bvrank  

<br/>

This file describes the structure bvrank.Rank and the bvrank package methods.   

<br/><br/>

---------------------------------------------  
<br/>  

## Structure: bvrank   
<br/>

The bvrank.Rank structure has the following fields:  
<br/>

- *Bitvector* (*bitvector.BitVector): Pointer to bitvector.BitVector, the underlying bitvector.  
<br/>

- *chcode* (uint8): This code indicates the most compact unsigned integer type in which the chunk ranks can be stored; 
    - If chcode is 0, chunk ranks are stored as uint8  
    - If chcode is 1, chunk ranks are stored as uint16  
    - If chcode is 2, chunk ranks are stored as uint32  
    - If chcode is 3, chunk ranks are stored as uint64    
<br/>

- *schcode* (uint8): This code indicates the most compact unsigned integer type in which the subchunk ranks can be stored; 
    - If schcode is 0, subchunk ranks are stored as uint8
    - If schcode is 1, subchunk ranks are stored as uint16
    - If schcode is 2, subchunk ranks are stored as uint32
    - If schcode is 3, subchunk ranks are stored as uint64  
<br/>

- *chranks\** ([]uint*): This slice (or array) stores the cumulative chunk ranks, where * is 8, 16, 32, or 64. Only one of chranks8, chranks16, chranks32, chranks64 will be non-empty (this matches the chcode).   
<br/>

- *schranks\** ([][]uint*): This slice of slices (or array of arrays) stores the relative subchunk ranks, where * is 8, 16, 32, or 64. Only one of schranks8, schranks16, schranks32, schranks64 will be non-empty (this matches the schcode).   
<br/><br/>


---------------------------------------------  
<br/>

## bitvector Methods
<br/>

The package bitvector has the following methods:  
<br/>

- **NewRankStruct**(*b* *bitvector.Bitvector) *Rank
   
    Creates new Rank structure over a bitvector.BitVector.  
    <br/>

    - Parameters:  
        - *b* (*bitvector.BitVector): Pointer to the bitvector over which to create the rank structure.

    <br/>

    - Returns:  
        - \*Rank: Pointer to the created bvrank.Rank struct created over the input bitvector *b*.  
    <br/>

    - Example call:  

        data := []byte{0, 0, 1, 0, 0, 0, 0, 1, 0, 1}
        var length uint64 = uint64(len(data))
        b := bitvector.NewBitVector(data, length)  
        r := bvrank.NewRankStruct(b)  

<br/>

- (*r* *Rank) **PopulateRanks**() *Rank  
  
    Fills out the Rank structure's rank values. Note that **NewRankStruct()** inherently calls **PopulateRanks**().  
    <br/>

    - No parameters; method on bvrank.Rank pointer.  
    <br/>

    - Returns: 
        - *Rank: Pointer to the updated bvrank.Rank struct with populated ranks.    
    <br/>

    - Example call:  

            r := r.PopulateRanks()
<br/>

- (*r* *Rank) **Rank1**(*i* uint64) uint64  
   
    Gets the rank1 of the bitvector at the input index *i*.  
    <br/>

     - Parameters:  
        - *i* (uint64): The index at which to find the rank.  
    <br/>

    - Returns: 
        - uint64: The rank of the bitvector at the input index *i*.  
    <br/>

    - Example call: 

            var i uint64 = 8
            ranki := r.Rank1(i)  
<br/>

- (*r* *Rank) **Overhead**() uint64  
   
    Gets the overhead of the current Rank struct.  
    <br/>

    - No parameters; method on bvrank.Rank pointer.  

    <br/>

    - Returns: 
        - uint64: The overhead (amount of bits) needed to store the current Rank structure.  
    <br/>

    - Example call: 

            overheadbits := r.Overhead()    
<br/>

- (*r* *Rank) **Save**(*filename* string)
    
    Saves the rank structure and the underlying bitvector structure to files.  
    <br/>

    - Parameters:  
        - *filename* (string):  Name of text file in which to save the Rank and BitVector structures. Note that Save() will change the end of the input filename (".txt") to "_rank.txt" and will save the underlying bitvector as the input filename with "_bitvector.txt" instead of ".txt".  
    <br/>

    - Does not return anything.  
    <br/>

    - Example call:  

            filename := "example_name.txt"
            r.Save(filename)  

        Note that this would create files "example_name_rank.txt" (which stores the rank struct without the bitvector) and "example_name_bitvector.txt" (which stores the bitvector struct).  

        The Load function takes this all into account when loading from a file.  
<br/>

- **Load**(*filename* string) *Rank  
   
    Loads the rank structure and underlying bitvector structure from files. Specifically, it loads the input filename with ".txt" replaced by "_rank.txt" and "_bitvector.txt" to load the rank and bitvector structs respectively.  
    <br/>

    - Parameters: 
        - *filename* (string): Name of text file from which to load the Rank and BitVector structures.  
    <br/>

    - Returns:  
        - *Rank: Pointer to the updated bvrank.Rank struct with populated ranks.   
    <br/>

    - Example call: 

            filename := "saved_bitvector.txt"
            r := bvrank.Load(filename)

         Note that this would load from files "example_name_rank.txt" (which stores the rank struct without the bitvector) and "example_name_bitvector.txt" (which stores the bitvector struct).    

<br/><br/>

### Supporting methods (not required by assignment but required to run implementation)  
<br/>

- (*r* *Rank) **ChunkLength**() uint64
    
    Gets the length of a chunk.  
    <br/>

    - No parameters; method on bvrank.Rank pointer.  

    <br/>

    - Returns: 
        - uint64: Length of a chunk in bits.  
    <br/>

    - Example call: 

            chunklength := r.ChunkLength()

<br/>

- (*r* *Rank) **NumberOfChunks**() uint64
    
    Gets the number of chunks the bitvector is broken into.  
    <br/>

    - No parameters; method on bvrank.Rank pointer.  

    <br/>

    - Returns: 
        - uint64: Number of chunks.  
    <br/>

    - Example call: 

            nchunks := r.NumberOfChunks()  

<br/>

- (*r* *Rank) **SubChunkLength**() uint64
    
    Gets the length of a subchunk.  
    <br/>

    - No parameters; method on bvrank.Rank pointer.  

    <br/>

    - Returns: 
        - uint64: Length of a subchunk in bits.  
    <br/>

    - Example call: 

            subchunklength := r.SubChunkLength()

<br/>

- (*r* *Rank) **NumberOfSubChunks**() uint64
    
    Gets the number of subchunks a chunk is broken into.  
    <br/>

    - No parameters; method on bvrank.Rank pointer.  

    <br/>

    - Returns: 
        - uint64: Number of subchunks per chunk.  
    <br/>

    - Example call: 

            nsubchunks := r.NumberOfSubChunks()  
            
<br/>

- (*r* *Rank) **RemainderChunkLength**() uint64
    
    Gets the length of the last (/remainder) chunk. Since the length of the bitvector may not always be a power of 2, we must account for the run-off in the last chunk.    
    <br/>

    - No parameters; method on bvrank.Rank pointer.  

    <br/>

    - Returns: 
        - uint64: Length of the last chunk in bits.  
    <br/>

    - Example call: 

            rchunklength := r.RemainderChunkLength()

<br/>

- (*r* *Rank) **RemainderNumberOfSubChunks**() uint64
    
    Gets the number of subchunks the last / remainder chunk is broken into.  
    <br/>

    - No parameters; method on bvrank.Rank pointer.  
    <br/>

    - Returns: 
        - uint64: Number of subchunks in the last chunk.  
    <br/>

    - Example call: 

            rnsubchunks := r.RemainderNumberOfSubChunks()  
            
<br/>

- (*r* *Rank) **accesschunkrank**(*i* uint64) uint64  
   
    Returns the chunk rank at the given index as a uint64.  
    <br/>

    - Parameters:  
        - *i* (uint64): Index of chunk.  

    <br/>

    - Returns: 
        - uint64: The cumulative rank at the start of the *i*th chunk.  
    <br/>

    - Example call: 

            var i uint64 = 0
            ind := r.accesschunkrank(i)    
<br/>

- (*r *Rank) **accesssubchunkrank**(*i* uint64, *j* uint64) uint64  
   
    Returns the subchunk rank at the given indices as a uint64.  
    <br/>

    - Parameters:  
        - *i* (uint64): Index of chunk.  

        - *j* (uint64): Index of subchunk.

    <br/>

    - Returns: 
        - uint64: The respective relative subchunk rank (the *j*th subchunk of chunk *i*).  
    <br/>

    - Example call: 

            var i uint64 = 0
            var j uint64 = 2
            ind := r.accesssubchunkrank(i, j)    
<br/>

- (*r *Rank) **ConvertToSmallest**(chranks64 []uint64, schranks [][]uint64) *Rank
   
    Converts the input arrays of cumulative chunk and relative subchunk ranks from uint64 to the smallest possible unsigned integers (uint8, uint16, or uint32).  
    <br/>

    - Parameters:  
        - *chranks* ([]uint64): Array (slice) of cumulative chunk ranks.  

        - *schranks* ([][]uint64): Array of arrays (slice of slices) of relative subchunk ranks.  

    <br/>

    - Does not return anything. Populates the corresponding fields of *r
    <br/>

    - Example call: 

            r = r.ConvertToSmallest(r.chranks64, r.schranks64)  

<br/>

- (*r *Rank) **ConvertToUint64**() ([]uint64, [][]uint64)
   
    Converts compressed cumulative chunk ranks and relative subchunk ranks (uint8, uint16, or uint32) to an array of uint64 and an array of arrays of uint64 values respectively. 
    <br/>

    - No parameters; method on bvrank.Rank pointer.  
    <br/>

    - Returns:  
        - []uint64: Array (slice) of cumulative chunk ranks.  

        - [][]uint64: Array of arrays (slice of slices) of relative subchunk ranks.  

    <br/>

    - Example call: 

            chranks64, schranks64 := r.ConvertToUint64()  

<br/>

- (*r *Rank) **GetRanks\***() ([]uint\*, [][]uint\*)
   
    Returns the respectively typed (uint8, uint16, uint32, uint64) cumulative chunk ranks and relative subchunk ranks of the current Rank structure (\* is one of 8, 16, 32, 64).
    <br/>

    - No parameters; method on bvrank.Rank pointer.  
    <br/>

    - Returns:  
        - []uint\*: Array (slice) of cumulative chunk ranks.  

        - [][]uint\*: Array of arrays (slice of slices) of relative subchunk ranks.  

    <br/>

    - Example call: 

            chranks8, schranks8 := r.GetRanks8()  
            chranks16, schranks16 := r.GetRanks16()  
            chranks32, schranks32 := r.GetRanks32()  
            chranks64, schranks64 := r.GetRanks64()  

<br/>

---------------------------------------------   
---------------------------------------------  
