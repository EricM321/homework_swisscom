import { Component, Input, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import * as _ from 'lodash';
import { map } from 'rxjs/internal/operators/map';


interface Feature {
  featureId: number;
  displayName: string;
  technicalName: string;
  expiresOn: string;
  description: string;
  inverted: boolean;
  active: boolean;
}

interface Features {
  features: Feature[]
}


@Component({
  selector: 'app-feature',
  templateUrl: './feature.component.html',
  styleUrls: ['./feature.component.scss']
})
export class FeatureComponent implements OnInit {

  // URL which returns list of JSON items (API end-point URL)
  private readonly URL = 'http://localhost:10000/api/v1/features';

  @Input()
  allFeatures: Observable<Feature[]>;

  constructor(private http: HttpClient) { 
    this.allFeatures = this.http.get<Feature[]>(this.URL).pipe(map((data: any) => _.values(data)));
    console.log(this.http.get<Features>(this.URL).pipe(map(data  => data)));
    console.log(this.allFeatures)
  }


  features: any[] = [];
  archivedFeatures: any[] = [];

  ngOnInit(){
      this.allFeatures.forEach(element => {
        if(element){
          this.features.push(element)
        }
        else {
          this.archivedFeatures.push(element)
        }
      });
  }


}
