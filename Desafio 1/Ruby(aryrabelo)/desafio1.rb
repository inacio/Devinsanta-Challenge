class Desafio1
  def self.perform
    (1..200).to_a.map do |num|
      number(num)
    end
  end

  def self.number(n)
    "#{n} #{show_santa?(n)}#{show_rem?(n)}".strip
  end

  def self.show_santa?(n)
    "Santa" if n % 5 == 0
  end

  def self.show_rem?(n)
    "rem" if n % 6 == 0
  end
end

puts Desafio1.perform