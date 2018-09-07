import { TestBed, inject } from '@angular/core/testing';

import { ExtService } from './ext.service';

describe('ExtService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ExtService]
    });
  });

  it('should be created', inject([ExtService], (service: ExtService) => {
    expect(service).toBeTruthy();
  }));
});
