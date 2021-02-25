import { Component, Inject } from '@angular/core';
import {MatDialog, MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import { Timestamp } from 'rxjs';

export interface CustomerData {
  name: string;
}

export interface FeatureData {
  displayName: string;
  technicalName: string;
  description: string;
  expiresOn: Timestamp<Date>;
  customers: string[];
}

/**
 * @title Dialog Overview
 */
@Component({
  selector: 'app-dialog-form',
  templateUrl: './dialog-form.component.html',
  styleUrls: ['./dialog-form.component.scss']
})
export class DialogFormComponent {

  name: string = "";
  displayName: string = "";
  technicalName: string = "";
  description: string = "";
  expiresOn: Date = new Date();
  customers: string[] = [];

  constructor(public dialog: MatDialog) {}

  openCustomerDialog(): void {
    const dialogRef = this.dialog.open(CustomerDialogForm, {
      width: '250px',
      data: {name: this.name}
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log('The dialog was closed');
      this.name = result;
    });
  }

  openFeatureDialog(): void {
    const dialogRef = this.dialog.open(FeatureDialogForm, {
      width: '500px',
      data: {
        displayName: this.displayName, 
        technicalName: this.technicalName, 
        description: this.description, 
        expiresOn: this.expiresOn, 
        customers: this.customers
      }
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log('The dialog was closed: ' + result);
      this.displayName = result;
    });
  }

}

@Component({
  selector: 'dialog-form',
  templateUrl: './customer-dialog-form.html',
})
export class CustomerDialogForm{

  constructor(
    public dialogRef: MatDialogRef<CustomerDialogForm>,
    @Inject(MAT_DIALOG_DATA) public data: CustomerData) {}

  onNoClick(): void {
    this.dialogRef.close();
  }

}

@Component({
  selector: 'dialog-form',
  templateUrl: './feature-dialog-form.html',
})
export class FeatureDialogForm{

  constructor(
    public dialogRef: MatDialogRef<FeatureDialogForm>,
    @Inject(MAT_DIALOG_DATA) public data: FeatureData) {}

  onNoClick(): void {
    this.dialogRef.close();
  }

}