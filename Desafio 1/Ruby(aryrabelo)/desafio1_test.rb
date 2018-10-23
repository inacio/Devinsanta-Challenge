require 'minitest/autorun'
require 'byebug'
require_relative 'desafio1'

class Desafio1Test < Minitest::Test
  def test_perform_1
    assert_equal Desafio1.perform.first, "1"
  end

  def test_perform_200
    assert_equal Desafio1.perform.last, "200 Santa"
  end


  def test_number_5
    assert_equal Desafio1.number(5), "5 Santa"
  end

  def test_number_6
    assert_equal Desafio1.number(6), "6 rem"
  end

  def test_number_30
    assert_equal Desafio1.number(30), "30 Santarem"
  end

  def test_number_1
    assert_equal Desafio1.number(1), "1"
  end
end
