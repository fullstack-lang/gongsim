import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongsimdatamodelComponent } from './gongsimdatamodel.component';

describe('GongsimdatamodelComponent', () => {
  let component: GongsimdatamodelComponent;
  let fixture: ComponentFixture<GongsimdatamodelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongsimdatamodelComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongsimdatamodelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
