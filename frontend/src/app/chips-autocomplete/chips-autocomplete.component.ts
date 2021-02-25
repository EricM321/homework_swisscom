import {COMMA, ENTER} from '@angular/cdk/keycodes';
import {Component, ElementRef, ViewChild} from '@angular/core';
import {FormControl} from '@angular/forms';
import {MatAutocompleteSelectedEvent, MatAutocomplete} from '@angular/material/autocomplete';
import {MatChipInputEvent} from '@angular/material/chips';
import {Observable} from 'rxjs';
import {map, startWith} from 'rxjs/operators';
import { allCustomers } from '../test.customers';


@Component({
  selector: 'app-chips-autocomplete',
  templateUrl: './chips-autocomplete.component.html',
  styleUrls: ['./chips-autocomplete.component.scss']
})
export class ChipsAutocompleteComponent {
  visible = true;
  selectable = true;
  removable = true;
  separatorKeysCodes: number[] = [ENTER, COMMA];
  customerCtrl = new FormControl();
  filteredCustomers: Observable<any[]>;
  customers: any[] = [];
  allCustomers: any[] = allCustomers;

  @ViewChild('customerInput')
  customerInput!: ElementRef<HTMLInputElement>;
  @ViewChild('auto') matAutocomplete!: MatAutocomplete;

  constructor() {
    this.filteredCustomers = this.customerCtrl.valueChanges.pipe(
        startWith(null),
        map((customer: string | null) => customer ? this._filter(customer) : this.allCustomers.slice()));
  }

  add(event: MatChipInputEvent): void {
    const input = event.input;
    const value = event.value;

    // Add our customer
    if ((value || '').trim()) {
      this.customers.push(value.trim());
    }

    // Reset the input value
    if (input) {
      input.value = '';
    }

    this.customerCtrl.setValue(null);
  }

  remove(customer: string): void {
    const index = this.customers.indexOf(customer);

    if (index >= 0) {
      this.customers.splice(index, 1);
    }
  }

  selected(event: MatAutocompleteSelectedEvent): void {
    this.customers.push(event.option.viewValue);
    this.customerInput.nativeElement.value = '';
    this.customerCtrl.setValue(null);
  }

  private _filter(value: string): string[] {
    const filterValue = value.toLowerCase();

    return this.allCustomers.filter(customer => customer.toLowerCase().indexOf(filterValue) === 0);
  }
}