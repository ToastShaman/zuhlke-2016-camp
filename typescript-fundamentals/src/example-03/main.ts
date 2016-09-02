interface Person {
  name: string
}

class Parent implements Person {
  constructor(public name: string) {}
}

class Father extends Parent {
  constructor(name: string) {
    super(name);
  }
}

class Mother extends Parent {
  constructor(name: string) {
    super(name);
  }
}

class Child extends Parent {
  constructor(name: string) {
    super(name);
  }
}

class FamilyTree<T extends Person> {
  private people: T[] = [];

  public add(person: T) {
    this.people.push(person);
  }

  public getPeople() {
    return this.people;
  }
}

const tree = new FamilyTree<Person>();
tree.add(new Father('Rene'));
tree.add(new Mother('Susanne'));
tree.add(new Mother('Kevin'));
tree.getPeople().forEach(p => console.log(`Person: ${p.name}`));
