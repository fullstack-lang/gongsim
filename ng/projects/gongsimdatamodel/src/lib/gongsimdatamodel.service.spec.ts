import { TestBed } from '@angular/core/testing';

import { GongsimdatamodelService } from './gongsimdatamodel.service';

describe('GongsimdatamodelService', () => {
  let service: GongsimdatamodelService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongsimdatamodelService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
