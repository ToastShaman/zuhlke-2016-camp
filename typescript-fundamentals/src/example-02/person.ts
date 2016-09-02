export enum Gender {
  MALE, FEMALE
}

export interface Person {
  firstname: string,
  middleName?: string,
  lastname: string,
  age: number,
  gender: Gender
}
