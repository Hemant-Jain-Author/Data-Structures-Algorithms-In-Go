package main

import (
	"fmt"
	"strconv"
)

type Polynomial struct {
     head*  Node
     tail*  Node
}

type Node struct {
     coeff  int
     pow  int
     next*  Node
}

func newNode( c  int,  p  int) * Node {
    var nd =  new(Node);
    nd.coeff = c;
    nd.pow = p;
    nd.next = nil;
    return nd;
}

func newPolynomial() * Polynomial {
    var poly * Polynomial =  &Polynomial{};
    poly.head = nil;
    poly.tail = nil;
    return poly;
}

func (this *Polynomial)assign(coeffs []int, pows []int,  size  int) *Polynomial {
    this.head = nil;
    this.tail = nil;
    var  temp*  Node = nil;
    for  i  := 0; i < size; i++ {
        temp = newNode(coeffs[i], pows[i]);
        if (this.head == nil) {
			this.head = temp;
        	this.tail = temp;
        } else {
            this.tail.next = temp;
            this.tail = temp;
        }
    }
    return this;
}

func (this *Polynomial) add( poly2*  Polynomial) *Polynomial {
    var  p1*  Node = this.head;
    var  p2*  Node = poly2.head;
    var  temp*  Node = nil;
    var  poly*  Polynomial = newPolynomial();
    for(p1 != nil || p2 != nil) {
        if (p1 == nil || (p2 != nil && p1.pow < p2.pow)) {
            temp = newNode(p2.coeff, p2.pow);
            p2 = p2.next;
        } else if (p2 == nil || p1.pow > p2.pow) {
            temp = newNode(p1.coeff, p1.pow);
            p1 = p1.next;
        } else if (p1.pow == p2.pow) {
            temp = newNode(p1.coeff + p2.coeff, p1.pow);
            p1 = p1.next;
            p2 = p2.next;
        }
        if (poly.head == nil) {
			poly.head = temp;
			poly.tail = temp;
        } else {
            poly.tail.next = temp;
            poly.tail = poly.tail.next;
        }
    }
    return poly;
}

func (this *Polynomial) print() {
    var  curr*  Node = this.head;
    for(curr != nil) {
        fmt.Print(strconv.Itoa(curr.coeff) + "x^" + strconv.Itoa(curr.pow));
        if (curr.next != nil) {
        fmt.Print(" + ");
        }
        curr = curr.next;
    }
    fmt.Println();
}

func main() {
    var c1  = []int{6, 5, 4}
    var p1  = []int{2, 1, 0}
    var  s1  int = len(c1);
    var  first*  Polynomial = newPolynomial().assign(c1, p1, s1);
    first.print();

    var c2  = []int{3, 2, 1}
    var p2  = []int{3, 1, 0}
    var  s2  int = len(c2);
    var  second*  Polynomial = newPolynomial().assign(c2, p2, s2);
    second.print();

    var  sum*  Polynomial = first.add(second);
    sum.print();
}

/*
6x^2 + 5x^1 + 4x^0
3x^3 + 2x^1 + 1x^0
3x^3 + 6x^2 + 7x^1 + 5x^0
*/
