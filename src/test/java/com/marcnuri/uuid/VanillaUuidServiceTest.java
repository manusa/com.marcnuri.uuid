/*
 * VanillaUuidServiceTest.java
 *
 * Created on 2019-07-07, 17:11
 */
package com.marcnuri.uuid;

import static org.hamcrest.Matchers.everyItem;
import static org.hamcrest.Matchers.hasSize;
import static org.hamcrest.Matchers.notNullValue;
import static org.junit.Assert.assertThat;

import java.util.Collection;
import java.util.UUID;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;

/**
 * Created by Marc Nuri <marc@marcnuri.com> on 2019-07-07.
 */
public class VanillaUuidServiceTest {

  private VanillaUuidService vanillaUuidService;

  @Before
  public void setUp() {
    vanillaUuidService = new VanillaUuidService();
  }

  @After
  public void tearDown() {
    vanillaUuidService = null;
  }

  @Test
  public void getUuid_na_shouldReturnUuid() {
    //When
    final UUID result = vanillaUuidService.getUuid();
    // Then
    assertThat(result, notNullValue());
  }


  @Test
  public void getUuid_quantityIs2_shouldReturn2UuidList() {
    // Given
    final int quantity = 2;
    //When
    final Collection<UUID> result = vanillaUuidService.getUuid(quantity);
    // Then
    assertThat(result, hasSize(2));
    assertThat(result, everyItem(notNullValue(UUID.class)));
  }
}
