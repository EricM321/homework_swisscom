import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import {MAT_FORM_FIELD_DEFAULT_OPTIONS} from '@angular/material/form-field';

import { MaterialModule } from './materials-module';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FeatureComponent } from './feature/feature.component';
import {DialogFormComponent, CustomerDialogForm, FeatureDialogForm} from './dialog-form/dialog-form.component';
import { ChipsAutocompleteComponent } from './chips-autocomplete/chips-autocomplete.component';

@NgModule({
  declarations: [
    AppComponent,
    FeatureComponent,
    DialogFormComponent,
    CustomerDialogForm,
    FeatureDialogForm,
    ChipsAutocompleteComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MaterialModule
  ],
  entryComponents: [AppComponent, FeatureComponent, DialogFormComponent, CustomerDialogForm, FeatureDialogForm],
  bootstrap: [AppComponent],
  providers: [{ provide: MAT_FORM_FIELD_DEFAULT_OPTIONS, useValue: { appearance: 'fill' } }]
})
export class AppModule { }
