import LinkedList from "./LinkedList.ts";

function testNewLinkedList() {
  const linkedList = new LinkedList<number>();

  console.assert(!linkedList.head, "expect LinkedList<> head to be blank");
  console.assert(!linkedList.tail, "expect LinkedList<> tail to be blank");
}

function testLinkedListPrepend() {
  const linkedList = new LinkedList<number>();
  linkedList.prepend(1);

  console.assert(!!linkedList.head, "expect LinkedList<> head to present");
  console.assert(!!linkedList.tail, "expect LinkedList<> tail to present");
}

function testLinkedListAppend() {
  const linkedList = new LinkedList<number>();
  linkedList.append(1);

  console.assert(!!linkedList.head, "expect LinkedList<> head to present");
  console.assert(!!linkedList.tail, "expect LinkedList<> tail to present");
}

function testLinkedListDeleteHead() {
  const linkedList = new LinkedList<number>();
  linkedList.append(1);
  linkedList.deleteHead();

  console.assert(!linkedList.head, "expect LinkedList<> head to be blank");
  console.assert(!linkedList.tail, "expect LinkedList<> tail to be blank");
}

function testLinkedListDeleteTail() {
  const linkedList = new LinkedList<number>();
  linkedList.prepend(1);
  linkedList.deleteTail();

  console.assert(!linkedList.head, "expect LinkedList<> head to be blank");
  console.assert(!linkedList.tail, "expect LinkedList<> tail to be blank");
}

testNewLinkedList();
testLinkedListPrepend();
testLinkedListAppend();
testLinkedListDeleteHead();
testLinkedListDeleteTail();
