# BLS Multisignature Implementation in Golang
This repo contains the implementation of BLS Multisignature scheme(https://medium.com/@manishas053/bls-signature-c52a5cbdcd92) using PBC library.

PBC Library is a free portable C library built on the GMP library that performs the mathematical operations underlying pairing based cryptosystems. It provides routines such as elliptic curve generation, elliptic curve arithmetic and pairing computations. Programs using the PBC library should include the file pbc.h and linked against the PBC library and GMP library.

To download PBC library, see https://crypto.stanford.edu/pbc/ and for GMP library, https://gmplib.org/

#For compiling and installing libraries 

 $ ./configure

 $ make

 $ sudo make install

It also requires the installation of PBC package (https://github.com/Nik-U/pbc). Package PBC provides structures for building pairing-based cryptosystems. It is a wrapper around the Pairing-Based Cryptography(PBC)  Library authored by Ben Lynn. During the build process, this package will attempt to include gmp.h and pbc/pbc.h, and then dynamically link to GMP and PBC libraries.
